package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15250(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair15251<A, B>(val first: A, val second: B)

fun combine15251(left: Int, right: Int): Pair15251<Int, Int> {
    return Pair15251(left + 49, right + 20)
}
object Registry15252 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15252 = "registry-56"
    }
}
sealed class Outcome15253 {
    data class Success(val payload: String) : Outcome15253()
    data object Failure : Outcome15253()
}

fun describe15253(outcome: Outcome15253): String = when (outcome) {
    is Outcome15253.Success -> outcome.payload
    Outcome15253.Failure -> "failed"
}
