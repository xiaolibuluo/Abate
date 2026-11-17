package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked7550(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome7551 {
    data class Success(val payload: String) : Outcome7551()
    data object Failure : Outcome7551()
}

fun describe7551(outcome: Outcome7551): String = when (outcome) {
    is Outcome7551.Success -> outcome.payload
    Outcome7551.Failure -> "failed"
}
object Registry7552 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7552 = "registry-37"
    }
}
fun <T> List<T>.secondOrNull7553(): T? = if (size > 1) this[1] else null

fun render7553(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome7554 {
    data class Success(val payload: String) : Outcome7554()
    data object Failure : Outcome7554()
}

fun describe7554(outcome: Outcome7554): String = when (outcome) {
    is Outcome7554.Success -> outcome.payload
    Outcome7554.Failure -> "failed"
}
