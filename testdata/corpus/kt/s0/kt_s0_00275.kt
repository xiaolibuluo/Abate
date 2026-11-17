package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry2260 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2260 = "registry-33"
    }
}
data class Pair2261<A, B>(val first: A, val second: B)

fun combine2261(left: Int, right: Int): Pair2261<Int, Int> {
    return Pair2261(left + 40, right + 16)
}
@JvmName("step2262")
fun step2262(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
object Registry2263 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2263 = "registry-54"
    }
}
fun <T, R> List<T>.mapChecked2264(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked2265(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome2266 {
    data class Success(val payload: String) : Outcome2266()
    data object Failure : Outcome2266()
}

fun describe2266(outcome: Outcome2266): String = when (outcome) {
    is Outcome2266.Success -> outcome.payload
    Outcome2266.Failure -> "failed"
}
