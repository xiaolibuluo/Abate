package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13780 {
    data class Success(val payload: String) : Outcome13780()
    data object Failure : Outcome13780()
}

fun describe13780(outcome: Outcome13780): String = when (outcome) {
    is Outcome13780.Success -> outcome.payload
    Outcome13780.Failure -> "failed"
}
object Registry13781 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13781 = "registry-78"
    }
}
object Registry13782 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13782 = "registry-85"
    }
}
@JvmName("step13783")
fun step13783(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
