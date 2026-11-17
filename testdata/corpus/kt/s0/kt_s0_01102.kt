package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9040 {
    data class Success(val payload: String) : Outcome9040()
    data object Failure : Outcome9040()
}

fun describe9040(outcome: Outcome9040): String = when (outcome) {
    is Outcome9040.Success -> outcome.payload
    Outcome9040.Failure -> "failed"
}
fun <T> List<T>.secondOrNull9041(): T? = if (size > 1) this[1] else null

fun render9041(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull9042(): T? = if (size > 1) this[1] else null

fun render9042(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry9043 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9043 = "registry-41"
    }
}
fun <T, R> List<T>.mapChecked9044(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry9045 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9045 = "registry-55"
    }
}
