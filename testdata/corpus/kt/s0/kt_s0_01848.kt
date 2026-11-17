package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15190 {
    data class Success(val payload: String) : Outcome15190()
    data object Failure : Outcome15190()
}

fun describe15190(outcome: Outcome15190): String = when (outcome) {
    is Outcome15190.Success -> outcome.payload
    Outcome15190.Failure -> "failed"
}
object Registry15191 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15191 = "registry-68"
    }
}
sealed class Outcome15192 {
    data class Success(val payload: String) : Outcome15192()
    data object Failure : Outcome15192()
}

fun describe15192(outcome: Outcome15192): String = when (outcome) {
    is Outcome15192.Success -> outcome.payload
    Outcome15192.Failure -> "failed"
}
