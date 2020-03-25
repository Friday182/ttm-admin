<template>
  <q-page class="flex-center">
    <div>
      <q-splitter
        v-model="splitterModel"
        style="height: 100%"
      >
        <template v-slot:before>
          <div
            class="q-pa-md"
            style="height: 100%;max-width: 350px"
          >
            <q-toolbar class="bg-primary text-white shadow-2">
              <q-toolbar-title>
                Contacts
              </q-toolbar-title>
            </q-toolbar>

            <q-list bordered>
              <q-item
                v-for="(contact, key) in contacts"
                :key="key"
                v-ripple
                class="q-my-sm"
                clickable
                :active="key == activeId"
                active-class="my-menu-link"
                @click="selectChat(key)"
              >
                <q-item-section avatar>
                  <q-avatar
                    color="primary"
                    text-color="white"
                  >
                    {{ contact.k_letter }}
                  </q-avatar>
                </q-item-section>
                <q-item-section>
                  <q-item-label>
                    {{ contact.k_name }}
                    <q-item-label
                      caption
                      lines="1"
                    >
                      {{ contact.k_role }}
                    </q-item-label>
                  </q-item-label>
                </q-item-section>
                <q-item-section side>
                  <q-icon
                    name="chat_bubble"
                    color="green"
                  />
                </q-item-section>
              </q-item>
              <q-separator />
            </q-list>
          </div>
        </template>
        <template v-slot:after>
          <q-separator />
          <div class="q-pa-md row justify-center">
            <q-card style="width: 80%; height: 50%; max-width: 1000px">
              <q-chat-message
                v-for="(msg, key) in showMsg"
                :key="key"
                :name="msg.sender"
                avatar="statics/icons/128.png"
                :text="msg.text"
                :sent="msg.send"
                :stamp="msg.timeStamp"
              />
              <q-btn
                v-if="showMsg.length===0"
                flat
                no-caps
                class="q-ma-xl q-pa-xl"
                size="lg"
                text-color="grey"
                label="You don't have any message ..."
              />
            </q-card>
          </div>
          <div class="q-pa-md q-gutter-sm">
            <q-editor
              v-model="editor"
              min-height="6rem"
            />
            <q-btn
              no-caps
              color="primary"
              label="Send"
              icon="mail"
              :disable="btnActive"
              @click="sendMsg"
            />
            <q-chip
              icon="error_outline"
              color="white"
              text-color="grey"
            >
              (Note: Message length between 20 to 2000 letters)
            </q-chip>
          </div>
        </template>
      </q-splitter>
      <alert-msg
        :active="alert"
        :msg="alertMsg"
        @cancelAlert="alertClose"
      />
    </div>
  </q-page>
</template>

<script type="text/javascript">
import { mapGetters, mapActions } from 'vuex'
import { ADD_MESSAGE_MUTATION } from '../../graphql/mutations'
import { READ_MESSAGE_QUERY } from '../../graphql/queries'

export default {
  name: 'Messages',
  components: {
    'alert-msg': require('components/common/AlertMsg.vue').default
  },

  data () {
    return {
      alert: false,
      alertMsg: '',
      splitterModel: 30,
      editor: '',
      activeId: 0,
      skipQuerySticker: false,
      skipQueryReadMsg: false,
      interHandler: 0,
      messages: [],
      allMessages: []
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),
    contacts: function () {
      return this.currentUser.contacts
    },
    btnActive: function () {
      if (this.editor.length > 20 && this.editor.length < 2000) {
        return false
      } else {
        return true
      }
    },
    curId: function () {
      return this.currentUser.id
    },
    curName: function () {
      return this.currentUser.name
    },
    showMsg: function () {
      let tmp = []
      for (let i = 0; i < this.messages.length; i++) {
        // console.log('tmp set', this.contacts[this.activeId].k_id, this.messages[i].txId, this.messages[i].rxId)
        if (this.contacts[this.activeId]) {
          if (this.contacts[this.activeId].k_id === this.messages[i].txId || this.contacts[this.activeId].k_id === this.messages[i].rxId) {
            tmp.push(this.messages[i])
          }
        }
      }
      // console.log('filter msg', JSON.stringify(tmp))
      return tmp
    }
  },
  mounted () {
    // this.interHandler = setInterval(this.refreshMsg, 10000)
    console.log('mentor contacts - ', JSON.stringify(this.currentUser.contacts))
  },
  beforeDestroy () {
    clearInterval(this.interHandler)
  },
  methods: {
    ...mapActions('student', ['updateStudent']),
    selectChat (id) {
      console.log('Chat with - ', id)
      this.activeId = id
    },
    sendMsg () {
      console.log('sending msg - ', this.editor.length)
      // send message
      this.$apollo
        .mutate({
          mutation: ADD_MESSAGE_MUTATION,
          variables: {
            txId: this.curId,
            rxId: this.contacts[this.activeId].k_id,
            sender: this.curName,
            receiver: this.contacts[this.activeId].k_name,
            text: this.editor
          }
        })
        .then(response => {
          if (response.data.addMessage.ok) {
            let tmpText = []
            tmpText.push(this.editor)
            this.messages.push({
              sender: 'Me',
              send: true,
              text: tmpText,
              txId: this.curId,
              rxId: this.contacts[this.activeId].k_id,
              timeStamp: 'a moment ago'
            })
            this.editor = ''
          } else {
            this.alertMsg = 'Add message failed'
            this.alert = true
          }
        })
        .catch(error => {
          this.alertMsg = 'Please contact your admistrator.'
          this.alert = true
          console.log(error)
        })
    },
    updateMsg (msg) {
      // console.log('Read msg - ', JSON.stringify(msg))
      let sender = ''
      let isSend = true
      let text = []
      this.messages = []
      for (let i = 0; i < msg.length; i++) {
        text = []
        if (msg[i].txId === this.curId) {
          sender = 'Me'
          isSend = true
        } else {
          sender = msg[i].sender
          isSend = false
        }
        text.push(msg[i].text)

        this.messages.push({
          sender: sender,
          send: isSend,
          text: text,
          txId: msg[i].txId,
          rxId: msg[i].rxId,
          timeStamp: String(msg[i].createDate).slice(0, 19)
        })
      }
      // console.log('messages - ', JSON.stringify(this.messages))
    },
    refreshMsg () {
      this.skipQueryReadMsg = false
      this.$apollo.queries.allMessages.refetch({
        id: this.curId,
        msgIdx: 0
      })
    },
    alertClose () {
      this.alert = false
      this.alertMsg = ''
    }
  },
  apollo: {
    allMessages: {
      query: READ_MESSAGE_QUERY,
      variables () {
        return {
          id: this.curId,
          msgIdx: 0
        }
      },
      error (error) {
        console.error('We\'ve got an error!', error)
      },
      skip () {
        return this.skipQueryReadMsg
      },
      result (data, key) {
        this.updateMsg(data.data.allMessages)
      }
    }
  }
}
</script>

<style lang="stylus" scoped>
.my-menu-link
  color: white
  background: #F2C037
</style>
