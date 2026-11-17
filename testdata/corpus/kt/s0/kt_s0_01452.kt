package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11990 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11990 = "registry-72"
    }
}
sealed class Outcome11991 {
    data class Success(val payload: String) : Outcome11991()
    data object Failure : Outcome11991()
}

fun describe11991(outcome: Outcome11991): String = when (outcome) {
    is Outcome11991.Success -> outcome.payload
    Outcome11991.Failure -> "failed"
}
object Registry11992 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11992 = "registry-86"
    }
}
object Registry11993 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11993 = "registry-93"
    }
}
sealed class Outcome11994 {
    data class Success(val payload: String) : Outcome11994()
    data object Failure : Outcome11994()
}

fun describe11994(outcome: Outcome11994): String = when (outcome) {
    is Outcome11994.Success -> outcome.payload
    Outcome11994.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11995(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
