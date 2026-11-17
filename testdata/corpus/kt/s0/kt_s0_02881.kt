package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23770 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
@JvmName("step23771")
fun step23771(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
class Counter23772 {
    var current: Int by Delegates.observable(74) { _, _, new -> println(new) }

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
object Registry23773 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23773 = "registry-81"
    }
}
@JvmName("step23774")
fun step23774(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
data class Pair23775<A, B>(val first: A, val second: B)

fun combine23775(left: Int, right: Int): Pair23775<Int, Int> {
    return Pair23775(left + 95, right + 50)
}
data class Pair23776<A, B>(val first: A, val second: B)

fun combine23776(left: Int, right: Int): Pair23776<Int, Int> {
    return Pair23776(left + 5, right + 61)
}
fun <T, R> List<T>.mapChecked23777(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
