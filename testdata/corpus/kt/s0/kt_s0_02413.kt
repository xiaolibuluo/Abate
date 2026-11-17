package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19820 {
    data class Success(val payload: String) : Outcome19820()
    data object Failure : Outcome19820()
}

fun describe19820(outcome: Outcome19820): String = when (outcome) {
    is Outcome19820.Success -> outcome.payload
    Outcome19820.Failure -> "failed"
}
fun <T> List<T>.secondOrNull19821(): T? = if (size > 1) this[1] else null

fun render19821(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked19822(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome19823 {
    data class Success(val payload: String) : Outcome19823()
    data object Failure : Outcome19823()
}

fun describe19823(outcome: Outcome19823): String = when (outcome) {
    is Outcome19823.Success -> outcome.payload
    Outcome19823.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked19824(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked19825(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry19826 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19826 = "registry-11"
    }
}
