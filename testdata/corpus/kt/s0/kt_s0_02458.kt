package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20200 {
    data class Success(val payload: String) : Outcome20200()
    data object Failure : Outcome20200()
}

fun describe20200(outcome: Outcome20200): String = when (outcome) {
    is Outcome20200.Success -> outcome.payload
    Outcome20200.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked20201(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry20202 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20202 = "registry-89"
    }
}
sealed class Outcome20203 {
    data class Success(val payload: String) : Outcome20203()
    data object Failure : Outcome20203()
}

fun describe20203(outcome: Outcome20203): String = when (outcome) {
    is Outcome20203.Success -> outcome.payload
    Outcome20203.Failure -> "failed"
}
fun <T> List<T>.secondOrNull20204(): T? = if (size > 1) this[1] else null

fun render20204(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
