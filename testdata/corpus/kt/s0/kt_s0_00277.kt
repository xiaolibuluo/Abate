package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2280 {
    data class Success(val payload: String) : Outcome2280()
    data object Failure : Outcome2280()
}

fun describe2280(outcome: Outcome2280): String = when (outcome) {
    is Outcome2280.Success -> outcome.payload
    Outcome2280.Failure -> "failed"
}
@JvmName("step2281")
fun step2281(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
@JvmName("step2282")
fun step2282(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
@JvmName("step2283")
fun step2283(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
sealed class Outcome2284 {
    data class Success(val payload: String) : Outcome2284()
    data object Failure : Outcome2284()
}

fun describe2284(outcome: Outcome2284): String = when (outcome) {
    is Outcome2284.Success -> outcome.payload
    Outcome2284.Failure -> "failed"
}
object Registry2285 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2285 = "registry-94"
    }
}
data class Pair2286<A, B>(val first: A, val second: B)

fun combine2286(left: Int, right: Int): Pair2286<Int, Int> {
    return Pair2286(left + 101, right + 24)
}
