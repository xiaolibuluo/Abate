package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome460 {
    data class Success(val payload: String) : Outcome460()
    data object Failure : Outcome460()
}

fun describe460(outcome: Outcome460): String = when (outcome) {
    is Outcome460.Success -> outcome.payload
    Outcome460.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked461(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome462 {
    data class Success(val payload: String) : Outcome462()
    data object Failure : Outcome462()
}

fun describe462(outcome: Outcome462): String = when (outcome) {
    is Outcome462.Success -> outcome.payload
    Outcome462.Failure -> "failed"
}
object Registry463 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG463 = "registry-42"
    }
}
fun <T> List<T>.secondOrNull464(): T? = if (size > 1) this[1] else null

fun render464(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair465<A, B>(val first: A, val second: B)

fun combine465(left: Int, right: Int): Pair465<Int, Int> {
    return Pair465(left + 56, right + 20)
}
object Registry466 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG466 = "registry-63"
    }
}
fun <T, R> List<T>.mapChecked467(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked468(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
