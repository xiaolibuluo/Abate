package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22840<A, B>(val first: A, val second: B)

fun combine22840(left: Int, right: Int): Pair22840<Int, Int> {
    return Pair22840(left + 15, right + 15)
}
class Counter22841 {
    var current: Int by Delegates.observable(22) { _, _, new -> println(new) }

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
sealed class Outcome22842 {
    data class Success(val payload: String) : Outcome22842()
    data object Failure : Outcome22842()
}

fun describe22842(outcome: Outcome22842): String = when (outcome) {
    is Outcome22842.Success -> outcome.payload
    Outcome22842.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked22843(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry22844 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22844 = "registry-43"
    }
}
fun <T> List<T>.secondOrNull22845(): T? = if (size > 1) this[1] else null

fun render22845(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step22846")
fun step22846(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
class Counter22847 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
