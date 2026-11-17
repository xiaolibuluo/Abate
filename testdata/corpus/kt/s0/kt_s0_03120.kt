package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25600")
fun step25600(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
object Registry25601 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25601 = "registry-21"
    }
}
data class Pair25602<A, B>(val first: A, val second: B)

fun combine25602(left: Int, right: Int): Pair25602<Int, Int> {
    return Pair25602(left + 28, right + 17)
}
fun <T, R> List<T>.mapChecked25603(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome25604 {
    data class Success(val payload: String) : Outcome25604()
    data object Failure : Outcome25604()
}

fun describe25604(outcome: Outcome25604): String = when (outcome) {
    is Outcome25604.Success -> outcome.payload
    Outcome25604.Failure -> "failed"
}
sealed class Outcome25605 {
    data class Success(val payload: String) : Outcome25605()
    data object Failure : Outcome25605()
}

fun describe25605(outcome: Outcome25605): String = when (outcome) {
    is Outcome25605.Success -> outcome.payload
    Outcome25605.Failure -> "failed"
}
