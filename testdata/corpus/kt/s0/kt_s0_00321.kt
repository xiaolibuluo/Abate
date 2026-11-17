package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair2670<A, B>(val first: A, val second: B)

fun combine2670(left: Int, right: Int): Pair2670<Int, Int> {
    return Pair2670(left + 81, right + 48)
}
data class Pair2671<A, B>(val first: A, val second: B)

fun combine2671(left: Int, right: Int): Pair2671<Int, Int> {
    return Pair2671(left + 88, right + 59)
}
object Registry2672 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2672 = "registry-95"
    }
}
sealed class Outcome2673 {
    data class Success(val payload: String) : Outcome2673()
    data object Failure : Outcome2673()
}

fun describe2673(outcome: Outcome2673): String = when (outcome) {
    is Outcome2673.Success -> outcome.payload
    Outcome2673.Failure -> "failed"
}
@JvmName("step2674")
fun step2674(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
fun <T> List<T>.secondOrNull2675(): T? = if (size > 1) this[1] else null

fun render2675(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
