package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3800<A, B>(val first: A, val second: B)

fun combine3800(left: Int, right: Int): Pair3800<Int, Int> {
    return Pair3800(left + 95, right + 46)
}
object Registry3801 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3801 = "registry-5"
    }
}
data class Pair3802<A, B>(val first: A, val second: B)

fun combine3802(left: Int, right: Int): Pair3802<Int, Int> {
    return Pair3802(left + 12, right + 68)
}
object Registry3803 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3803 = "registry-19"
    }
}
fun <T, R> List<T>.mapChecked3804(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome3805 {
    data class Success(val payload: String) : Outcome3805()
    data object Failure : Outcome3805()
}

fun describe3805(outcome: Outcome3805): String = when (outcome) {
    is Outcome3805.Success -> outcome.payload
    Outcome3805.Failure -> "failed"
}
@JvmName("step3806")
fun step3806(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
data class Pair3807<A, B>(val first: A, val second: B)

fun combine3807(left: Int, right: Int): Pair3807<Int, Int> {
    return Pair3807(left + 47, right + 62)
}
sealed class Outcome3808 {
    data class Success(val payload: String) : Outcome3808()
    data object Failure : Outcome3808()
}

fun describe3808(outcome: Outcome3808): String = when (outcome) {
    is Outcome3808.Success -> outcome.payload
    Outcome3808.Failure -> "failed"
}
