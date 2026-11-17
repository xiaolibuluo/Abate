package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6520 {
    data class Success(val payload: String) : Outcome6520()
    data object Failure : Outcome6520()
}

fun describe6520(outcome: Outcome6520): String = when (outcome) {
    is Outcome6520.Success -> outcome.payload
    Outcome6520.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6521(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6522<A, B>(val first: A, val second: B)

fun combine6522(left: Int, right: Int): Pair6522<Int, Int> {
    return Pair6522(left + 56, right + 20)
}
@JvmName("step6523")
fun step6523(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
object Registry6524 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6524 = "registry-70"
    }
}
