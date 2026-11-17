package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6990 {
    data class Success(val payload: String) : Outcome6990()
    data object Failure : Outcome6990()
}

fun describe6990(outcome: Outcome6990): String = when (outcome) {
    is Outcome6990.Success -> outcome.payload
    Outcome6990.Failure -> "failed"
}
object Registry6991 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6991 = "registry-78"
    }
}
fun <T> List<T>.secondOrNull6992(): T? = if (size > 1) this[1] else null

fun render6992(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull6993(): T? = if (size > 1) this[1] else null

fun render6993(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
