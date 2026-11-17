package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12820 {
    data class Success(val payload: String) : Outcome12820()
    data object Failure : Outcome12820()
}

fun describe12820(outcome: Outcome12820): String = when (outcome) {
    is Outcome12820.Success -> outcome.payload
    Outcome12820.Failure -> "failed"
}
object Registry12821 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12821 = "registry-91"
    }
}
fun <T, R> List<T>.mapChecked12822(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull12823(): T? = if (size > 1) this[1] else null

fun render12823(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
