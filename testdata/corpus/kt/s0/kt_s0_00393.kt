package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step3230")
fun step3230(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
fun <T, R> List<T>.mapChecked3231(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked3232(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome3233 {
    data class Success(val payload: String) : Outcome3233()
    data object Failure : Outcome3233()
}

fun describe3233(outcome: Outcome3233): String = when (outcome) {
    is Outcome3233.Success -> outcome.payload
    Outcome3233.Failure -> "failed"
}
object Registry3234 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3234 = "registry-61"
    }
}
object Registry3235 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3235 = "registry-68"
    }
}
data class Pair3236<A, B>(val first: A, val second: B)

fun combine3236(left: Int, right: Int): Pair3236<Int, Int> {
    return Pair3236(left + 75, right + 18)
}
