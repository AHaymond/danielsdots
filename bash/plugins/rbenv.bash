if [ -n "$(which rbenv)" ]; then
  export RBENV_ROOT="$HOME/.rbenv"
  export PATH="$HOME/.rbenv/bin:$PATH"
  eval "$(rbenv init -)"
fi
