package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked23170(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome23171 {
    data class Success(val payload: String) : Outcome23171()
    data object Failure : Outcome23171()
}

fun describe23171(outcome: Outcome23171): String = when (outcome) {
    is Outcome23171.Success -> outcome.payload
    Outcome23171.Failure -> "failed"
}
fun <T> List<T>.secondOrNull23172(): T? = if (size > 1) this[1] else null

fun render23172(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked23173(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry23174 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23174 = "registry-84"
    }
}
object Registry23175 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23175 = "registry-91"
    }
}
