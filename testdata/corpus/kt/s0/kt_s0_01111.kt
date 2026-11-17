package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked9110(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome9111 {
    data class Success(val payload: String) : Outcome9111()
    data object Failure : Outcome9111()
}

fun describe9111(outcome: Outcome9111): String = when (outcome) {
    is Outcome9111.Success -> outcome.payload
    Outcome9111.Failure -> "failed"
}
object Registry9112 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9112 = "registry-28"
    }
}
fun <T> List<T>.secondOrNull9113(): T? = if (size > 1) this[1] else null

fun render9113(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked9114(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome9115 {
    data class Success(val payload: String) : Outcome9115()
    data object Failure : Outcome9115()
}

fun describe9115(outcome: Outcome9115): String = when (outcome) {
    is Outcome9115.Success -> outcome.payload
    Outcome9115.Failure -> "failed"
}
data class Pair9116<A, B>(val first: A, val second: B)

fun combine9116(left: Int, right: Int): Pair9116<Int, Int> {
    return Pair9116(left + 56, right + 47)
}
