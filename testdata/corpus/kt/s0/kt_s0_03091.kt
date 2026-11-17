package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry25380 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25380 = "registry-19"
    }
}
object Registry25381 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25381 = "registry-26"
    }
}
sealed class Outcome25382 {
    data class Success(val payload: String) : Outcome25382()
    data object Failure : Outcome25382()
}

fun describe25382(outcome: Outcome25382): String = when (outcome) {
    is Outcome25382.Success -> outcome.payload
    Outcome25382.Failure -> "failed"
}
