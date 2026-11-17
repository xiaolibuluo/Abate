package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome1690 {
    data class Success(val payload: String) : Outcome1690()
    data object Failure : Outcome1690()
}

fun describe1690(outcome: Outcome1690): String = when (outcome) {
    is Outcome1690.Success -> outcome.payload
    Outcome1690.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked1691(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair1692<A, B>(val first: A, val second: B)

fun combine1692(left: Int, right: Int): Pair1692<Int, Int> {
    return Pair1692(left + 82, right + 55)
}
object Registry1693 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1693 = "registry-89"
    }
}
fun <T, R> List<T>.mapChecked1694(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
