package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20640 {
    data class Success(val payload: String) : Outcome20640()
    data object Failure : Outcome20640()
}

fun describe20640(outcome: Outcome20640): String = when (outcome) {
    is Outcome20640.Success -> outcome.payload
    Outcome20640.Failure -> "failed"
}
sealed class Outcome20641 {
    data class Success(val payload: String) : Outcome20641()
    data object Failure : Outcome20641()
}

fun describe20641(outcome: Outcome20641): String = when (outcome) {
    is Outcome20641.Success -> outcome.payload
    Outcome20641.Failure -> "failed"
}
sealed class Outcome20642 {
    data class Success(val payload: String) : Outcome20642()
    data object Failure : Outcome20642()
}

fun describe20642(outcome: Outcome20642): String = when (outcome) {
    is Outcome20642.Success -> outcome.payload
    Outcome20642.Failure -> "failed"
}
sealed class Outcome20643 {
    data class Success(val payload: String) : Outcome20643()
    data object Failure : Outcome20643()
}

fun describe20643(outcome: Outcome20643): String = when (outcome) {
    is Outcome20643.Success -> outcome.payload
    Outcome20643.Failure -> "failed"
}
