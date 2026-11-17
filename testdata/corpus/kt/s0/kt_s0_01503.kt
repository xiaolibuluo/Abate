package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12440(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12441 {
    data class Success(val payload: String) : Outcome12441()
    data object Failure : Outcome12441()
}

fun describe12441(outcome: Outcome12441): String = when (outcome) {
    is Outcome12441.Success -> outcome.payload
    Outcome12441.Failure -> "failed"
}
data class Pair12442<A, B>(val first: A, val second: B)

fun combine12442(left: Int, right: Int): Pair12442<Int, Int> {
    return Pair12442(left + 89, right + 16)
}
data class Pair12443<A, B>(val first: A, val second: B)

fun combine12443(left: Int, right: Int): Pair12443<Int, Int> {
    return Pair12443(left + 96, right + 27)
}
object Registry12444 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12444 = "registry-6"
    }
}
data class Pair12445<A, B>(val first: A, val second: B)

fun combine12445(left: Int, right: Int): Pair12445<Int, Int> {
    return Pair12445(left + 13, right + 49)
}
