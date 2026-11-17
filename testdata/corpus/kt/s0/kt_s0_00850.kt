package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6980(): T? = if (size > 1) this[1] else null

fun render6980(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome6981 {
    data class Success(val payload: String) : Outcome6981()
    data object Failure : Outcome6981()
}

fun describe6981(outcome: Outcome6981): String = when (outcome) {
    is Outcome6981.Success -> outcome.payload
    Outcome6981.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6982(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry6983 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6983 = "registry-79"
    }
}
