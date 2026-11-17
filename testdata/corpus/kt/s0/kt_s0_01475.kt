package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12210 {
    data class Success(val payload: String) : Outcome12210()
    data object Failure : Outcome12210()
}

fun describe12210(outcome: Outcome12210): String = when (outcome) {
    is Outcome12210.Success -> outcome.payload
    Outcome12210.Failure -> "failed"
}
