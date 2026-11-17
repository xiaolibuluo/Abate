package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked22680(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull22681(): T? = if (size > 1) this[1] else null

fun render22681(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked22682(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome22683 {
    data class Success(val payload: String) : Outcome22683()
    data object Failure : Outcome22683()
}

fun describe22683(outcome: Outcome22683): String = when (outcome) {
    is Outcome22683.Success -> outcome.payload
    Outcome22683.Failure -> "failed"
}
object Registry22684 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22684 = "registry-29"
    }
}
sealed class Outcome22685 {
    data class Success(val payload: String) : Outcome22685()
    data object Failure : Outcome22685()
}

fun describe22685(outcome: Outcome22685): String = when (outcome) {
    is Outcome22685.Success -> outcome.payload
    Outcome22685.Failure -> "failed"
}
