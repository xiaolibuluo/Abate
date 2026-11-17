package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12450(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12451 {
    data class Success(val payload: String) : Outcome12451()
    data object Failure : Outcome12451()
}

fun describe12451(outcome: Outcome12451): String = when (outcome) {
    is Outcome12451.Success -> outcome.payload
    Outcome12451.Failure -> "failed"
}
sealed class Outcome12452 {
    data class Success(val payload: String) : Outcome12452()
    data object Failure : Outcome12452()
}

fun describe12452(outcome: Outcome12452): String = when (outcome) {
    is Outcome12452.Success -> outcome.payload
    Outcome12452.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12453(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked12454(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry12455 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12455 = "registry-26"
    }
}
