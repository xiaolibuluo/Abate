package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4190(): T? = if (size > 1) this[1] else null

fun render4190(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome4191 {
    data class Success(val payload: String) : Outcome4191()
    data object Failure : Outcome4191()
}

fun describe4191(outcome: Outcome4191): String = when (outcome) {
    is Outcome4191.Success -> outcome.payload
    Outcome4191.Failure -> "failed"
}
object Registry4192 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4192 = "registry-34"
    }
}
fun <T, R> List<T>.mapChecked4193(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry4194 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4194 = "registry-48"
    }
}
object Registry4195 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4195 = "registry-55"
    }
}
