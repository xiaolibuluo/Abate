package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull12500(): T? = if (size > 1) this[1] else null

fun render12500(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked12501(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry12502 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12502 = "registry-70"
    }
}
fun <T> List<T>.secondOrNull12503(): T? = if (size > 1) this[1] else null

fun render12503(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull12504(): T? = if (size > 1) this[1] else null

fun render12504(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull12505(): T? = if (size > 1) this[1] else null

fun render12505(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome12506 {
    data class Success(val payload: String) : Outcome12506()
    data object Failure : Outcome12506()
}

fun describe12506(outcome: Outcome12506): String = when (outcome) {
    is Outcome12506.Success -> outcome.payload
    Outcome12506.Failure -> "failed"
}
