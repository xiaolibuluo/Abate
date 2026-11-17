package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked4690(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step4691")
fun step4691(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
fun <T, R> List<T>.mapChecked4692(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry4693 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4693 = "registry-12"
    }
}
fun <T, R> List<T>.mapChecked4694(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome4695 {
    data class Success(val payload: String) : Outcome4695()
    data object Failure : Outcome4695()
}

fun describe4695(outcome: Outcome4695): String = when (outcome) {
    is Outcome4695.Success -> outcome.payload
    Outcome4695.Failure -> "failed"
}
object Registry4696 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4696 = "registry-33"
    }
}
