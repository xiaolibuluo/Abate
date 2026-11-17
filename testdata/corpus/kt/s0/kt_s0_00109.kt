package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome870 {
    data class Success(val payload: String) : Outcome870()
    data object Failure : Outcome870()
}

fun describe870(outcome: Outcome870): String = when (outcome) {
    is Outcome870.Success -> outcome.payload
    Outcome870.Failure -> "failed"
}
fun <T> List<T>.secondOrNull871(): T? = if (size > 1) this[1] else null

fun render871(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked872(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull873(): T? = if (size > 1) this[1] else null

fun render873(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome874 {
    data class Success(val payload: String) : Outcome874()
    data object Failure : Outcome874()
}

fun describe874(outcome: Outcome874): String = when (outcome) {
    is Outcome874.Success -> outcome.payload
    Outcome874.Failure -> "failed"
}
object Registry875 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG875 = "registry-7"
    }
}
fun <T> List<T>.secondOrNull876(): T? = if (size > 1) this[1] else null

fun render876(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
