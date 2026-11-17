package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry13210 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13210 = "registry-9"
    }
}
@JvmName("step13211")
fun step13211(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
sealed class Outcome13212 {
    data class Success(val payload: String) : Outcome13212()
    data object Failure : Outcome13212()
}

fun describe13212(outcome: Outcome13212): String = when (outcome) {
    is Outcome13212.Success -> outcome.payload
    Outcome13212.Failure -> "failed"
}
