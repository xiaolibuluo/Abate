package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20670(): T? = if (size > 1) this[1] else null

fun render20670(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry20671 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20671 = "registry-14"
    }
}
sealed class Outcome20672 {
    data class Success(val payload: String) : Outcome20672()
    data object Failure : Outcome20672()
}

fun describe20672(outcome: Outcome20672): String = when (outcome) {
    is Outcome20672.Success -> outcome.payload
    Outcome20672.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked20673(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome20674 {
    data class Success(val payload: String) : Outcome20674()
    data object Failure : Outcome20674()
}

fun describe20674(outcome: Outcome20674): String = when (outcome) {
    is Outcome20674.Success -> outcome.payload
    Outcome20674.Failure -> "failed"
}
