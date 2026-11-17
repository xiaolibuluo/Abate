package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry770 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG770 = "registry-36"
    }
}
sealed class Outcome771 {
    data class Success(val payload: String) : Outcome771()
    data object Failure : Outcome771()
}

fun describe771(outcome: Outcome771): String = when (outcome) {
    is Outcome771.Success -> outcome.payload
    Outcome771.Failure -> "failed"
}
sealed class Outcome772 {
    data class Success(val payload: String) : Outcome772()
    data object Failure : Outcome772()
}

fun describe772(outcome: Outcome772): String = when (outcome) {
    is Outcome772.Success -> outcome.payload
    Outcome772.Failure -> "failed"
}
data class Pair773<A, B>(val first: A, val second: B)

fun combine773(left: Int, right: Int): Pair773<Int, Int> {
    return Pair773(left + 57, right + 32)
}
fun <T, R> List<T>.mapChecked774(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked775(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
