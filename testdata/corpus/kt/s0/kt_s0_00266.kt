package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2180 {
    data class Success(val payload: String) : Outcome2180()
    data object Failure : Outcome2180()
}

fun describe2180(outcome: Outcome2180): String = when (outcome) {
    is Outcome2180.Success -> outcome.payload
    Outcome2180.Failure -> "failed"
}
fun <T> List<T>.secondOrNull2181(): T? = if (size > 1) this[1] else null

fun render2181(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry2182 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2182 = "registry-40"
    }
}
fun <T, R> List<T>.mapChecked2183(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull2184(): T? = if (size > 1) this[1] else null

fun render2184(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
