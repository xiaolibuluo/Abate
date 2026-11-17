package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7400 {
    data class Success(val payload: String) : Outcome7400()
    data object Failure : Outcome7400()
}

fun describe7400(outcome: Outcome7400): String = when (outcome) {
    is Outcome7400.Success -> outcome.payload
    Outcome7400.Failure -> "failed"
}
object Registry7401 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7401 = "registry-29"
    }
}
object Registry7402 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7402 = "registry-36"
    }
}
