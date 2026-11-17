package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5540(): T? = if (size > 1) this[1] else null

fun render5540(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull5541(): T? = if (size > 1) this[1] else null

fun render5541(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5542 {
    data class Success(val payload: String) : Outcome5542()
    data object Failure : Outcome5542()
}

fun describe5542(outcome: Outcome5542): String = when (outcome) {
    is Outcome5542.Success -> outcome.payload
    Outcome5542.Failure -> "failed"
}
object Registry5543 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5543 = "registry-50"
    }
}
