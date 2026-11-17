package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4120 {
    data class Success(val payload: String) : Outcome4120()
    data object Failure : Outcome4120()
}

fun describe4120(outcome: Outcome4120): String = when (outcome) {
    is Outcome4120.Success -> outcome.payload
    Outcome4120.Failure -> "failed"
}
object Registry4121 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4121 = "registry-33"
    }
}
@JvmName("step4122")
fun step4122(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
data class Pair4123<A, B>(val first: A, val second: B)

fun combine4123(left: Int, right: Int): Pair4123<Int, Int> {
    return Pair4123(left + 47, right + 59)
}
@JvmName("step4124")
fun step4124(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
@JvmName("step4125")
fun step4125(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
