package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull25280(): T? = if (size > 1) this[1] else null

fun render25280(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked25281(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair25282<A, B>(val first: A, val second: B)

fun combine25282(left: Int, right: Int): Pair25282<Int, Int> {
    return Pair25282(left + 97, right + 37)
}
fun <T, R> List<T>.mapChecked25283(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25284(): T? = if (size > 1) this[1] else null

fun render25284(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter25285 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
object Registry25286 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25286 = "registry-28"
    }
}
sealed class Outcome25287 {
    data class Success(val payload: String) : Outcome25287()
    data object Failure : Outcome25287()
}

fun describe25287(outcome: Outcome25287): String = when (outcome) {
    is Outcome25287.Success -> outcome.payload
    Outcome25287.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25288(): T? = if (size > 1) this[1] else null

fun render25288(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
