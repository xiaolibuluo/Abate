package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17350 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17350 = "registry-56"
    }
}
sealed class Outcome17351 {
    data class Success(val payload: String) : Outcome17351()
    data object Failure : Outcome17351()
}

fun describe17351(outcome: Outcome17351): String = when (outcome) {
    is Outcome17351.Success -> outcome.payload
    Outcome17351.Failure -> "failed"
}
sealed class Outcome17352 {
    data class Success(val payload: String) : Outcome17352()
    data object Failure : Outcome17352()
}

fun describe17352(outcome: Outcome17352): String = when (outcome) {
    is Outcome17352.Success -> outcome.payload
    Outcome17352.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked17353(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome17354 {
    data class Success(val payload: String) : Outcome17354()
    data object Failure : Outcome17354()
}

fun describe17354(outcome: Outcome17354): String = when (outcome) {
    is Outcome17354.Success -> outcome.payload
    Outcome17354.Failure -> "failed"
}
data class Pair17355<A, B>(val first: A, val second: B)

fun combine17355(left: Int, right: Int): Pair17355<Int, Int> {
    return Pair17355(left + 91, right + 9)
}
