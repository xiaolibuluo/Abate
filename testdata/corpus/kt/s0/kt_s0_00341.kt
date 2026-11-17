package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step2840")
fun step2840(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
@JvmName("step2841")
fun step2841(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
fun <T> List<T>.secondOrNull2842(): T? = if (size > 1) this[1] else null

fun render2842(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome2843 {
    data class Success(val payload: String) : Outcome2843()
    data object Failure : Outcome2843()
}

fun describe2843(outcome: Outcome2843): String = when (outcome) {
    is Outcome2843.Success -> outcome.payload
    Outcome2843.Failure -> "failed"
}
@JvmName("step2844")
fun step2844(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
data class Pair2845<A, B>(val first: A, val second: B)

fun combine2845(left: Int, right: Int): Pair2845<Int, Int> {
    return Pair2845(left + 46, right + 39)
}
sealed class Outcome2846 {
    data class Success(val payload: String) : Outcome2846()
    data object Failure : Outcome2846()
}

fun describe2846(outcome: Outcome2846): String = when (outcome) {
    is Outcome2846.Success -> outcome.payload
    Outcome2846.Failure -> "failed"
}
object Registry2847 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2847 = "registry-60"
    }
}
fun <T> List<T>.secondOrNull2848(): T? = if (size > 1) this[1] else null

fun render2848(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
