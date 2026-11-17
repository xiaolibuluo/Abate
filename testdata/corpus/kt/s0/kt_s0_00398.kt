package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry3270 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3270 = "registry-85"
    }
}
data class Pair3271<A, B>(val first: A, val second: B)

fun combine3271(left: Int, right: Int): Pair3271<Int, Int> {
    return Pair3271(left + 92, right + 52)
}
fun <T> List<T>.secondOrNull3272(): T? = if (size > 1) this[1] else null

fun render3272(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull3273(): T? = if (size > 1) this[1] else null

fun render3273(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome3274 {
    data class Success(val payload: String) : Outcome3274()
    data object Failure : Outcome3274()
}

fun describe3274(outcome: Outcome3274): String = when (outcome) {
    is Outcome3274.Success -> outcome.payload
    Outcome3274.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked3275(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter3276 {
    var current: Int by Delegates.observable(30) { _, _, new -> println(new) }

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
