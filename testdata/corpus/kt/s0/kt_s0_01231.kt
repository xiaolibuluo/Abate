package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter10170 {
    var current: Int by Delegates.observable(34) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked10171(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull10172(): T? = if (size > 1) this[1] else null

fun render10172(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry10173 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10173 = "registry-55"
    }
}
data class Pair10174<A, B>(val first: A, val second: B)

fun combine10174(left: Int, right: Int): Pair10174<Int, Int> {
    return Pair10174(left + 62, right + 35)
}
fun <T, R> List<T>.mapChecked10175(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair10176<A, B>(val first: A, val second: B)

fun combine10176(left: Int, right: Int): Pair10176<Int, Int> {
    return Pair10176(left + 76, right + 57)
}
sealed class Outcome10177 {
    data class Success(val payload: String) : Outcome10177()
    data object Failure : Outcome10177()
}

fun describe10177(outcome: Outcome10177): String = when (outcome) {
    is Outcome10177.Success -> outcome.payload
    Outcome10177.Failure -> "failed"
}
