package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18310 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18310 = "registry-43"
    }
}
object Registry18311 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18311 = "registry-50"
    }
}
fun <T, R> List<T>.mapChecked18312(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18313 {
    data class Success(val payload: String) : Outcome18313()
    data object Failure : Outcome18313()
}

fun describe18313(outcome: Outcome18313): String = when (outcome) {
    is Outcome18313.Success -> outcome.payload
    Outcome18313.Failure -> "failed"
}
fun <T> List<T>.secondOrNull18314(): T? = if (size > 1) this[1] else null

fun render18314(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull18315(): T? = if (size > 1) this[1] else null

fun render18315(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
