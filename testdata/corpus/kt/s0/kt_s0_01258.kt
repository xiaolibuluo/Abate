package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10390 {
    data class Success(val payload: String) : Outcome10390()
    data object Failure : Outcome10390()
}

fun describe10390(outcome: Outcome10390): String = when (outcome) {
    is Outcome10390.Success -> outcome.payload
    Outcome10390.Failure -> "failed"
}
sealed class Outcome10391 {
    data class Success(val payload: String) : Outcome10391()
    data object Failure : Outcome10391()
}

fun describe10391(outcome: Outcome10391): String = when (outcome) {
    is Outcome10391.Success -> outcome.payload
    Outcome10391.Failure -> "failed"
}
sealed class Outcome10392 {
    data class Success(val payload: String) : Outcome10392()
    data object Failure : Outcome10392()
}

fun describe10392(outcome: Outcome10392): String = when (outcome) {
    is Outcome10392.Success -> outcome.payload
    Outcome10392.Failure -> "failed"
}
