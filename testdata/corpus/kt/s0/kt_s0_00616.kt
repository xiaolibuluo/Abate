package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5020 {
    data class Success(val payload: String) : Outcome5020()
    data object Failure : Outcome5020()
}

fun describe5020(outcome: Outcome5020): String = when (outcome) {
    is Outcome5020.Success -> outcome.payload
    Outcome5020.Failure -> "failed"
}
sealed class Outcome5021 {
    data class Success(val payload: String) : Outcome5021()
    data object Failure : Outcome5021()
}

fun describe5021(outcome: Outcome5021): String = when (outcome) {
    is Outcome5021.Success -> outcome.payload
    Outcome5021.Failure -> "failed"
}
sealed class Outcome5022 {
    data class Success(val payload: String) : Outcome5022()
    data object Failure : Outcome5022()
}

fun describe5022(outcome: Outcome5022): String = when (outcome) {
    is Outcome5022.Success -> outcome.payload
    Outcome5022.Failure -> "failed"
}
