package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18400<A, B>(val first: A, val second: B)

fun combine18400(left: Int, right: Int): Pair18400<Int, Int> {
    return Pair18400(left + 63, right + 18)
}
data class Pair18401<A, B>(val first: A, val second: B)

fun combine18401(left: Int, right: Int): Pair18401<Int, Int> {
    return Pair18401(left + 70, right + 29)
}
class Counter18402 {
    var current: Int by Delegates.observable(77) { _, _, new -> println(new) }

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
data class Pair18403<A, B>(val first: A, val second: B)

fun combine18403(left: Int, right: Int): Pair18403<Int, Int> {
    return Pair18403(left + 84, right + 51)
}
object Registry18404 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18404 = "registry-91"
    }
}
fun <T, R> List<T>.mapChecked18405(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter18406 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
class Counter18407 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
