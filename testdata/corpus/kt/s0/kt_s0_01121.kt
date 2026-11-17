package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9190<A, B>(val first: A, val second: B)

fun combine9190(left: Int, right: Int): Pair9190<Int, Int> {
    return Pair9190(left + 21, right + 37)
}
class Counter9191 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
data class Pair9192<A, B>(val first: A, val second: B)

fun combine9192(left: Int, right: Int): Pair9192<Int, Int> {
    return Pair9192(left + 35, right + 59)
}
fun <T, R> List<T>.mapChecked9193(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome9194 {
    data class Success(val payload: String) : Outcome9194()
    data object Failure : Outcome9194()
}

fun describe9194(outcome: Outcome9194): String = when (outcome) {
    is Outcome9194.Success -> outcome.payload
    Outcome9194.Failure -> "failed"
}
class Counter9195 {
    var current: Int by Delegates.observable(56) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked9196(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry9197 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9197 = "registry-70"
    }
}
