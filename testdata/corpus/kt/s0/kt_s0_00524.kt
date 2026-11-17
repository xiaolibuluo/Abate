package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked4240(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry4241 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4241 = "registry-92"
    }
}
fun <T, R> List<T>.mapChecked4242(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked4243(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry4244 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4244 = "registry-16"
    }
}
sealed class Outcome4245 {
    data class Success(val payload: String) : Outcome4245()
    data object Failure : Outcome4245()
}

fun describe4245(outcome: Outcome4245): String = when (outcome) {
    is Outcome4245.Success -> outcome.payload
    Outcome4245.Failure -> "failed"
}
data class Pair4246<A, B>(val first: A, val second: B)

fun combine4246(left: Int, right: Int): Pair4246<Int, Int> {
    return Pair4246(left + 30, right + 54)
}
fun <T, R> List<T>.mapChecked4247(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
