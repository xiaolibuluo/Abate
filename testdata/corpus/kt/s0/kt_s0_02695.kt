package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22270<A, B>(val first: A, val second: B)

fun combine22270(left: Int, right: Int): Pair22270<Int, Int> {
    return Pair22270(left + 50, right + 43)
}
class Counter22271 {
    var current: Int by Delegates.observable(57) { _, _, new -> println(new) }

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
sealed class Outcome22272 {
    data class Success(val payload: String) : Outcome22272()
    data object Failure : Outcome22272()
}

fun describe22272(outcome: Outcome22272): String = when (outcome) {
    is Outcome22272.Success -> outcome.payload
    Outcome22272.Failure -> "failed"
}
object Registry22273 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22273 = "registry-71"
    }
}
class Counter22274 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
@JvmName("step22275")
fun step22275(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
fun <T, R> List<T>.mapChecked22276(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
