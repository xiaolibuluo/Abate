package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4780(): T? = if (size > 1) this[1] else null

fun render4780(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4781(): T? = if (size > 1) this[1] else null

fun render4781(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4782(): T? = if (size > 1) this[1] else null

fun render4782(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4783 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4783 = "registry-32"
    }
}
sealed class Outcome4784 {
    data class Success(val payload: String) : Outcome4784()
    data object Failure : Outcome4784()
}

fun describe4784(outcome: Outcome4784): String = when (outcome) {
    is Outcome4784.Success -> outcome.payload
    Outcome4784.Failure -> "failed"
}
fun <T> List<T>.secondOrNull4785(): T? = if (size > 1) this[1] else null

fun render4785(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked4786(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked4787(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
