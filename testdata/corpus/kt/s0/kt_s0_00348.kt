package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry2900 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2900 = "registry-89"
    }
}
sealed class Outcome2901 {
    data class Success(val payload: String) : Outcome2901()
    data object Failure : Outcome2901()
}

fun describe2901(outcome: Outcome2901): String = when (outcome) {
    is Outcome2901.Success -> outcome.payload
    Outcome2901.Failure -> "failed"
}
data class Pair2902<A, B>(val first: A, val second: B)

fun combine2902(left: Int, right: Int): Pair2902<Int, Int> {
    return Pair2902(left + 6, right + 48)
}
fun <T, R> List<T>.mapChecked2903(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked2904(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked2905(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair2906<A, B>(val first: A, val second: B)

fun combine2906(left: Int, right: Int): Pair2906<Int, Int> {
    return Pair2906(left + 34, right + 31)
}
object Registry2907 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2907 = "registry-41"
    }
}
object Registry2908 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2908 = "registry-48"
    }
}
fun <T, R> List<T>.mapChecked2909(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
