package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19250 {
    data class Success(val payload: String) : Outcome19250()
    data object Failure : Outcome19250()
}

fun describe19250(outcome: Outcome19250): String = when (outcome) {
    is Outcome19250.Success -> outcome.payload
    Outcome19250.Failure -> "failed"
}
object Registry19251 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19251 = "registry-11"
    }
}
sealed class Outcome19252 {
    data class Success(val payload: String) : Outcome19252()
    data object Failure : Outcome19252()
}

fun describe19252(outcome: Outcome19252): String = when (outcome) {
    is Outcome19252.Success -> outcome.payload
    Outcome19252.Failure -> "failed"
}
@JvmName("step19253")
fun step19253(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
data class Pair19254<A, B>(val first: A, val second: B)

fun combine19254(left: Int, right: Int): Pair19254<Int, Int> {
    return Pair19254(left + 32, right + 47)
}
sealed class Outcome19255 {
    data class Success(val payload: String) : Outcome19255()
    data object Failure : Outcome19255()
}

fun describe19255(outcome: Outcome19255): String = when (outcome) {
    is Outcome19255.Success -> outcome.payload
    Outcome19255.Failure -> "failed"
}
