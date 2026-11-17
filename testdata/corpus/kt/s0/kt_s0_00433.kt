package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3540 {
    data class Success(val payload: String) : Outcome3540()
    data object Failure : Outcome3540()
}

fun describe3540(outcome: Outcome3540): String = when (outcome) {
    is Outcome3540.Success -> outcome.payload
    Outcome3540.Failure -> "failed"
}
object Registry3541 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3541 = "registry-55"
    }
}
sealed class Outcome3542 {
    data class Success(val payload: String) : Outcome3542()
    data object Failure : Outcome3542()
}

fun describe3542(outcome: Outcome3542): String = when (outcome) {
    is Outcome3542.Success -> outcome.payload
    Outcome3542.Failure -> "failed"
}
