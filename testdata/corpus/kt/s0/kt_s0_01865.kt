package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15330 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15330 = "registry-49"
    }
}
fun <T, R> List<T>.mapChecked15331(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked15332(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair15333<A, B>(val first: A, val second: B)

fun combine15333(left: Int, right: Int): Pair15333<Int, Int> {
    return Pair15333(left + 70, right + 37)
}
sealed class Outcome15334 {
    data class Success(val payload: String) : Outcome15334()
    data object Failure : Outcome15334()
}

fun describe15334(outcome: Outcome15334): String = when (outcome) {
    is Outcome15334.Success -> outcome.payload
    Outcome15334.Failure -> "failed"
}
